#include <iostream>
#include <algorithm>
#include <queue>
#include <unordered_map>
#include <unordered_set>

#include "graph.h"

using pq_type = std::priority_queue <Vertex, std::vector<Vertex>, std::greater<>>;

class PriorityQ {
public:
    // changes the priority (vertex value) of queue element
    void change_priority(uint idx, uint value) {
        auto iter = m_vertices.find(idx);
        if (iter != m_vertices.end())
            if (iter->second->get_value() <= value)
                return;
        iter->second->set_expired();

        auto vtx = Vertex{idx, value};
        m_pq.push(vtx);  // to avoid copy, try use emplace.
        m_vertices.insert({idx, &vtx});
    }

    // removes and returns the top element of the queue.
    Vertex min_priority(bool remove = true) {
        auto top = m_pq.top();
        while (top.is_expired() && !m_pq.empty()) {
            m_pq.pop();
            top = m_pq.top();
        }
        if (remove)
            m_pq.pop();
        return top;
    }

    // insert queue_element into queue
    void insert(uint idx, uint value) {
        change_priority(idx, value);
    }

    // returns the top element of the queue.
    Vertex top() {
        return min_priority(false);
    }

    // return the number of queue_elements.
    inline size_t size() {
        return m_pq.size();
    }

    inline bool empty() {
        return m_pq.empty();
    }

private:
    pq_type m_pq;
    std::unordered_map<uint, Vertex *> m_vertices;
};

class ShortestPath {
public:
    explicit ShortestPath(Graph &dij_graph) : m_graph(&dij_graph) {
        // pre-allocation
        m_closed_set.reserve(vertex_number);
    }

    void Solve() {
        auto open_set = PriorityQ();
        std::unordered_set <uint> closed_set;
        auto vn = m_graph->vertex_n();

        // starts from vertex-0
        open_set.insert(0, 0);

        while (!open_set.empty()) {
            auto top = open_set.min_priority();
            auto top_i = top.get_index();
            closed_set.insert(top_i);

            auto sub_arr = *(m_graph->sub_array(top_i));
            for (int i = 0; i < vn; i++) {
                if (i == top_i || closed_set.find(i) != closed_set.end())
                    continue;
                // add adjacency vertex to open set
                open_set.insert(i, sub_arr[i]);
            }
        }

        auto vn = m_graph->vertex_n();
        for (int i = 1; i < vn; ++i) {
            for (int j = 1; j < vn; ++j)
                if (!m_closed_set[i] && m_graph->adjacent(0, i))
                    m_open_set.push_back(j);
        }
    }

    void Print() const {
        m_graph->Print();
    }

private:
    // graph
    Graph *m_graph;
    // closed set, contains vertices already expanded
    std::vector<bool> m_closed_set;
    //
    PriorityQ m_pq;
};

int main() {
    std::cout << "Hello, World!" << std::endl;

    auto graph = Graph(0.1, 10);
    auto dij = ShortestPath(graph);
    dij.Solve();
    dij.Print();

    return 0;
}

