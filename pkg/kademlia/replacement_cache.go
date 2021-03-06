// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package kademlia

import (
	proto "storj.io/storj/protos/overlay"
	"storj.io/storj/storage"
)


func (rt *RoutingTable) addToReplacementCache(kadBucketID storage.Key, node *proto.Node) {	
	bucketID := string(kadBucketID)
	nodes := rt.replacementCache[bucketID]
	nodes = append(nodes, node)
	if len(nodes) > rt.rcBucketSize {
		copy(nodes, nodes[1:])
		nodes = nodes[:len(nodes)-1]
	}
	rt.replacementCache[bucketID] = nodes
}
