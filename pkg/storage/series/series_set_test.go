// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/querier/series/series_set_test.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package series

import (
	"testing"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/stretchr/testify/require"
)

func TestConcreteSeriesSet(t *testing.T) {
	series1 := &ConcreteSeries{
		labels:  labels.FromStrings("foo", "bar"),
		samples: []model.SamplePair{{Value: 1, Timestamp: 2}},
	}
	series2 := &ConcreteSeries{
		labels:  labels.FromStrings("foo", "baz"),
		samples: []model.SamplePair{{Value: 3, Timestamp: 4}},
	}
	c := NewConcreteSeriesSet([]storage.Series{series2, series1})
	require.True(t, c.Next())
	require.Equal(t, series1, c.At())
	require.True(t, c.Next())
	require.Equal(t, series2, c.At())
	require.False(t, c.Next())
}

func TestMatrixToSeriesSetSortsMetricLabels(t *testing.T) {
	matrix := model.Matrix{
		{
			Metric: model.Metric{
				model.MetricNameLabel: "testmetric",
				"e":                   "f",
				"a":                   "b",
				"g":                   "h",
				"c":                   "d",
			},
			Values: []model.SamplePair{{Timestamp: 0, Value: 0}},
		},
	}
	ss := MatrixToSeriesSet(matrix)
	require.True(t, ss.Next())
	require.NoError(t, ss.Err())

	l := ss.At().Labels()
	require.Equal(t, labels.FromStrings(model.MetricNameLabel, "testmetric", "a", "b", "c", "d", "e", "f", "g", "h"), l)
}
