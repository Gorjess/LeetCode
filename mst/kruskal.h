//
// Created by gorjess on 5/4/2022.
//

#ifndef MST_KRUSKAL_H
#define MST_KRUSKAL_H

#include <iostream>
#include <algorithm>
#include <utility>
#include "graph-theory/common/alias.h"
#include "graph-theory/common/graph.h"
#include "graph-theory/common/priority_q.h"


// { weight: {from, to} }
using edges_type = PriorityQ <std::pair<uint, std::pair < uint, uint>>>;


class KruskalSolution {

public:
    // Constructor
    KruskalSolution(Graph &graph) {
        // sort edges by weight.
        for (uint i = 0; i < graph.vertex_n(); ++i)
            for (uint j = 0; j < graph.vertex_n(); ++j) {
                auto pr = std::make_pair(graph.xy_weight(i, j), std::make_pair(i, j));
                m_edges.insert(pr);
            }
    }

    // determine whether two vertices are part of the same tree


private:
    edges_type m_edges;
};


#endif //MST_KRUSKAL_H