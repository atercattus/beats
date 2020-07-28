// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package panw

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "panw", asset.ModuleFieldsPri, AssetPanw); err != nil {
		panic(err)
	}
}

// AssetPanw returns asset data.
// This is the base64 encoded gzipped contents of module/panw.
func AssetPanw() string {
	return "eJzMmM9u4zYQxu95irm5BeLcesmhQNpFgABpauw26HFBUyOJNcXRDkdxtU+/IC3LWouOtE42iG6STH6/+fNRpJewwfYaauW2FwBixGJ/l6HXbGox5K7h9wsAgL8oayxCTgwrZQlurBA8oGyJNx5+Wd08LP/+9OsFQG7QZv46DlqCU9Vh2nBJW+M1FExN3T1JiIXrNs4DOVMFUmKcA6pIcdX9aCgFR3rkD49Tqs9Jf6dPHOUTQe9iBkuFvxqOHWENyLix6FG+l+rwNthuibOjd88xAsCDqhAoj4xhcpBSCVRKdIkZSGk8ePTekLtKAnlqWGOSZ5SuaZouaUKA/wu6LGIJ1UuLT2g7MaD1f6jl6mh0Km1D0q/kjjmncjeDOFyfdlhBoKv3qawNeYwT5FyNkve6UL3KD5A5ddxg8GxRZxKtyMvy4eaffRlVljF6fwkm3z8Kb42HGjknrjAbM56u8yCzKcB9ACdezuAfR3C3SgH2qwhxKo97EEuueD2UIHaASVo1Qy/GqTDvG/l1oPjuTPthwPa+nDske4/2HVZ16OHh85cZedLKE2ae6aETMaVdPcPXk85+GdcMi7vd3uKN7N2pnWntWqv6s0l55BU8tFJ6gwJa1dIwwt2H6B8FUjKqU1mEF5poTkNrqqrGGWnToc8Jf2YKwvXnXi1mwNJ2WSpf9pvS0GG/LaWpD3viE42VG/tWm7wgdWZLheDO65c/jFPcQp+dfafsaDw6CbxrBOWUbb9iui7rNsbyr7HZreEwjp+MxtRyki5yMvcN2zdKfcP2zMxrJVgQtz/HzbexX2M9Hj/eh6+NLHxkf/x432unV+0wdl+QyzjmCTkzWoBcvI0VVi4D45MToJESGRaVskYbavziEhYFq3arGBeXQAyLNTpTuMWUiSxtx7Z/weHtLuwOnLLgmgrZaDAZOjG5QY5djEqX4/1C+hyHXxp0Gj+7plojJxkT37UJwHsqAJ1wOySLJ0zjwTjNWKETzDp5McraRB0fnfnS4CEkS0VEmoipW+wZnzmknpX30HXEu84JUnM+LkdQr9kGif8Wjhoh2vwH6MLNT+Xrs5YiGwIpffK8ModlOYK5iROCqA26nqD3yLcAAAD//yC/rB0="
}