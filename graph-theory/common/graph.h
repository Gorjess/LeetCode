//
// Created by gorjess on 5/4/2022.
//

#ifndef GRAPH_THEORY_COMMON_GRAPH_H
#define GRAPH_THEORY_COMMON_GRAPH_H

#include <iostream>
#include <array>
#include <algorithm>
#include <queue>

#include "alias.h"
#include "random_engine.h"

/*
 * An Undirected graph.
 *
 * An adjacency matrix is used to represent the graph,
 * for it is always faster than adjacency list.
 * Dynamicly choose the underlying ADT by density would be more elegant,
 * but it requires more work to make the code applies to both ADT.
 * */
class Graph {
public:
    Graph() = delete;

    Graph(const Graph &) = delete;

    Graph(float density, int edge_range) {
        // convert density to a integer between 0 and 10
        auto factor = density <= 0.0f ? 0.1f : density <= 1.0f ? density : 1.0f;
        m_density = static_cast<int>(factor * 10);

        // initialize graph according to density.
        auto edge_random_engine = RandomEngine<int>(1, edge_range);
        for (int i = 0; i < vertex_number; i++)
            for (int j = 0; j < vertex_number; j++)
                if (i == j)
                    m_graph[i][j] = 0;
                else if (!m_graph[j][i])
                    m_graph[i][j] =
                    m_graph[j][i] = m_re_density.Gen() < m_density ? edge_random_engine.Gen() : 0;
    }

    // "[[nodiscard]]" suggests that you should always use the returned value of a getter
    [[nodiscard]] int vertex_n() const { return m_vertex_n; }

    [[nodiscard]] int edge_n() const { return m_edge_n; }

    // tests whether there is an edge from node x to node y.
    bool adjacent(uint x, uint y) { return m_graph[x][y] == 0; }

    // return weight at (x, y)
    [[nodiscard]] uint xy_weight(uint x, uint y) { return m_graph[x][y]; }

    // lists all nodes y such that there is an edge from x to y.
    std::vector<uint> neighbors(uint x) {
        std::vector<uint> nbrs;
        auto vn = vertex_n();
        if (x >= vn)
            return nbrs;
        for (int i = 0; i < vn; ++i)
            if (m_graph[x][i])
                nbrs.push_back(i);
        return nbrs;
    }

    // list sub-array by index i
    std::array<uint, vertex_number> *sub_array(uint i) {
        if (i < 0 || i >= vertex_n())
            return nullptr;
        return &m_graph[i];
    }

    // adds to m_graph the edge from x to y, if it is not there.
    void add(uint x, uint y, uint weight) {
        auto vn = vertex_n();
        if (x >= vn || y >= vn)
            return;
        if (m_graph[x][y])
            m_graph[x][y] = weight;
    }

    // removes the edge from x to y, if it is there.
    void delete_edge(uint x, uint y) {
        auto vn = vertex_n();
        if (x >= vn || y >= vn)
            return;
        if (m_graph[x][y])
            m_graph[x][y] = 0;
    }

    void Print() const {
        std::cout << "==================Graph: ==================\n";
        for (auto &sub_arr: m_graph) {
            for (auto &entry: sub_arr)
                std::cout << entry << " ";
            std::cout << std::endl;
        }
        std::cout << "===========================================\n";
    }

private:
    // vertex number
    int m_vertex_n = vertex_number;
    // edge number
    int m_edge_n = 0;
    // density random engine
    int m_density;
    RandomEngine<int> m_re_density = RandomEngine<int>(10);
    // adjacency matrix describing connectivity and edge-value.
    std::array<std::array<uint, vertex_number>, vertex_number> m_graph{};
};

class Vertex {
public:
    Vertex(uint i, uint v) : m_index(i), m_value(v), m_expired(false) {}

    void set_value(uint v) { m_value = v; }

    void set_expired() { m_expired = true; }

    [[nodiscard]] uint get_index() const { return m_index; }

    [[nodiscard]] uint get_value() const { return m_value; }

    [[nodiscard]] uint is_expired() const { return m_expired; }

    inline bool operator>(const Vertex &v) const { return m_value > v.m_value; }

private:
    uint m_index;
    uint m_value;
    // indicates if this vertex is expired, that is,
    // the value of vertex at index m_index has changed.
    bool m_expired;
};

#endif //GRAPH_THEORY_COMMON_GRAPH_H
