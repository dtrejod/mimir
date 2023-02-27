// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/ruler/manager_test.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package ruler

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/test"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/notifier"
	"github.com/prometheus/prometheus/rules"
	promRules "github.com/prometheus/prometheus/rules"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/grafana/mimir/pkg/ruler/rulespb"
)

func TestSyncRuleGroups(t *testing.T) {
	dir := t.TempDir()

	m, err := NewDefaultMultiTenantManager(Config{RulePath: dir}, factory, nil, log.NewNopLogger(), nil)
	require.NoError(t, err)

	const (
		user1  = "testUser1"
		user2 = "testUser2"
	)

	userRules := map[string]rulespb.RuleGroupList{
		user: {
			&rulespb.RuleGroupDesc{
				Name:      "group1",
				Namespace: "ns",
				Interval:  1 * time.Minute,
				User:      user,
			},
		},
		user1: {
			&rulespb.RuleGroupDesc{
				Name:      "group2",
				Namespace: "ns",
				Interval:  1 * time.Minute,
				User:      user1,
			},
		},
	}
	m.SyncRuleGroups(context.Background(), userRules)

	mgr1 := getManager(m, user)
	require.NotNil(t, mgr)
	test.Poll(t, 1*time.Second, true, func() interface{} {
		return mgr.(*mockRulesManager).running.Load()
	})

	mgr2 := getManager(m, user2)
	require.NotNil(t, mgr1)
	test.Poll(t, 1*time.Second, true, func() interface{} {
		return mgr1.(*mockRulesManager).running.Load()
	})

	// Verify that user rule groups are now cached locally.
	{
		users, err := m.mapper.users()
		require.NoError(t, err)
		require.Contains(t, users, user)
		require.Contains(t, users, user1)
		require.Len(t, users, 2)
	}

	// Passing empty map / nil stops all managers.
	m.SyncRuleGroups(context.Background(), nil)
	require.Nil(t, getManager(m, user))

	// Make sure old manager was stopped.
	test.Poll(t, 1*time.Second, false, func() interface{} {
		return mgr.(*mockRulesManager).running.Load()
	})

	test.Poll(t, 1*time.Second, false, func() interface{} {
		return mgr1.(*mockRulesManager).running.Load()
	})

	// Verify that local rule groups were removed.
	{
		users, err := m.mapper.users()
		require.NoError(t, err)
		require.Equal(t, []string(nil), users)
	}

	// Resync same rules as before. Previously this didn't restart the manager.
	m.SyncRuleGroups(context.Background(), userRules)

	newMgr := getManager(m, user)
	require.NotNil(t, newMgr)
	require.True(t, mgr != newMgr)

	test.Poll(t, 1*time.Second, true, func() interface{} {
		return newMgr.(*mockRulesManager).running.Load()
	})

	// Verify that user rule groups are cached locally again.
	{
		users, err := m.mapper.users()
		require.NoError(t, err)
		require.Contains(t, users, user)
	}

	m.Stop()

	test.Poll(t, 1*time.Second, false, func() interface{} {
		return newMgr.(*mockRulesManager).running.Load()
	})
}

func TestSyncRulesToManager(t *testing.T) {
	dir := t.TempDir()

	m, err := NewDefaultMultiTenantManager(Config{RulePath: dir}, factory, nil, log.NewNopLogger(), nil)
	require.NoError(t, err)

	const (
		user = "testUser"
	)

	userRules := make(map[string]rulespb.RuleGroupList)

	for i := 0; i < 11; i++ {
		userID := user + strconv.Itoa(i)
		userRules[userID] = rulespb.RuleGroupList{
			&rulespb.RuleGroupDesc{
				Name:      "group1",
				Namespace: "ns",
				Interval:  1 * time.Minute,
				User:      userID,
			}}
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			userID := user + strconv.Itoa(idx)
			m.syncRulesToManager(context.Background(), userID, userRules[userID])
		}(i)
	}
	wg.Wait()

	ruleManagers := []RulesManager{}
	for i := 0; i < 11; i++ {
		userID := user + strconv.Itoa(i)
		mgr := getManager(m, userID)
		ruleManagers = append(ruleManagers, mgr)
		require.NotNil(t, mgr)
		test.Poll(t, 1*time.Second, true, func() interface{} {
			return mgr.(*mockRulesManager).running.Load()
		})

		users, err := m.mapper.users()
		require.NoError(t, err)
		require.Contains(t, users, userID)
	}

	m.Stop()

	for _, rm := range ruleManagers {
		test.Poll(t, 1*time.Second, false, func() interface{} {
			return rm.(*mockRulesManager).running.Load()
		})
	}
}

func getManager(m *DefaultMultiTenantManager, user string) RulesManager {
	m.userManagerMtx.RLock()
	defer m.userManagerMtx.RUnlock()

	return m.userManagers[user]
}

func factory(_ context.Context, _ string, _ *notifier.Manager, _ log.Logger, _ prometheus.Registerer) RulesManager {
	return &mockRulesManager{done: make(chan struct{})}
}

type mockRulesManager struct {
	running atomic.Bool
	done    chan struct{}
}

func (m *mockRulesManager) Run() {
	m.running.Store(true)
	<-m.done
}

func (m *mockRulesManager) Stop() {
	m.running.Store(false)
	close(m.done)
}

func (m *mockRulesManager) Update(interval time.Duration, files []string, externalLabels labels.Labels, externalURL string, ruleGroupPostProcessFunc rules.RuleGroupPostProcessFunc) error {
	return nil
}

func (m *mockRulesManager) RuleGroups() []*promRules.Group {
	return nil
}
