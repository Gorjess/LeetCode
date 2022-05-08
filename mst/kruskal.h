//
// Created by gorjess on 5/4/2022.
//

#ifndef MST_KRUSKAL_H
#define MST_KRUSKAL_H

#include <iostream>
#include <algorithm>
#include <utility>
#include <unordered_map>
#include <unordered_set>
#include "../graph-theory/common/alias.h"
#include "../graph-theory/common/graph.h"


// { weight: {from, to} }
using edges_type = std::vector<std::array<uint, 3>>;


typedef struct _DisjointSet {
    std::vector<uint> parents;
    std::vector<uint> rank;
    uint m_max_index;

    explicit _DisjointSet(uint vertex_num) : m_max_index(vertex_num) {
        parents.reserve(vertex_num);
        rank.reserve(vertex_num);
        for (uint i = 0; i < vertex_num; ++i) {
            parents[i] = i;
            rank[i] = 1;
        }
    }

    // return the parent vertex index of node_index.
    // parent is updated after searching.
    uint find_parent(uint idx) {
        if (parents[idx] == idx)
            return idx;
        return parents[idx] = find_parent(parents[idx]);
    }

    // unite x and y to one same set.
    void unite(uint x, uint y) {
        auto px = find_parent(x);
        auto py = find_parent(y);

        if (px == py)
            return;

        if (rank[x] < rank[y]) {
            parents[px] = py;
            rank[py] += rank[px];
        } else {
            parents[py] = px;
            rank[px] += rank[py];
        }
    }

} DisjointSet;


class KruskalSolution {

public:
    // Constructor
    KruskalSolution(Graph &graph) : m_vertex_n(graph.vertex_n()) {
        m_sets.reserve(m_vertex_n);
        // sort edges by weight.
        for (uint i = 0; i < m_vertex_n; ++i) {
            m_sets[i] = i;  // make set
            for (uint j = 0; j < graph.vertex_n(); ++j) {
                m_edges.push_back({graph.xy_weight(i, j), i, j});
            }
        }
    }

    // return total weight of a minimum spanning tree
    uint mst_solution() {
        // sort edges by their weight
        std::sort(m_edges.begin(), m_edges.end());

        DisjointSet ds(m_vertex_n);
        uint weight_sum = 0;
        for (auto& edge: m_edges) {
            auto w = edge[0];
            auto x = edge[1];
            auto y = edge[2];
            if (ds.find_parent(x) != ds.find_parent(y)) {
                ds.unite(x, y);
                weight_sum += w;
            }
        }

        return weight_sum;
    }


private:
    uint m_vertex_n;
    edges_type m_edges;
    std::vector <uint> m_sets;
};


#endif //MST_KRUSKAL_H