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
#include "graph-theory/common/alias.h"
#include "graph-theory/common/graph.h"
#include "graph-theory/common/priority_q.h"


// { weight: {from, to} }
using edges_type = PriorityQ <std::pair<uint, std::pair < uint, uint>>>;

typedef struct _KKComponent {
    uint vtx_index;  
    _KKComponent* parent;
    std::unordered_map<uint, _KKComponent*> leaves;

    bool append_leaf(uint index, _KKComponent* leaf) {
        auto iter = leaves.find(index);
        if (iter == leaves.end()) 
            return false;
        iter->second->add_leaf(index, leaf);
        return true;
    }

    inline void add_leaf(uint index, _KKComponent* leaf) {
        leaves.insert(index, leaf);
    }
} KKComponent;

typedef struct _KKTree {
    KKComponent* root;
    std::unordered_set<uint> indexes;

    // return whether index is contained in tree
    bool contain(const uint& index) {
        return indexes.find(index) != indexes.end();
    }
} KKTree;

class KruskalSolution {

public:
    KruskalSolution(Graph &graph) {
        // sort edges by weight.
        for (uint i = 0; i < graph.vertex_n(); ++i)
            for (uint j = 0; j < graph.vertex_n(); ++j) {
                auto pr = std::make_pair(graph.xy_weight(i, j), std::make_pair(i, j));
                m_edges.insert(pr);
            }
    }

private:
    edges_type m_edges;

};


#endif //MST_KRUSKAL_H