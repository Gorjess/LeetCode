#include <iostream>
#include <queue>
#include <unordered_map>

#include "graph.h"

using pq_type = std::priority_queue<Vertex, std::vector<Vertex>, std::greater<>>;

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
        m_pq.push(vtx);  // todo: try use emplace.
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
        auto vn = m_graph->vertex_n();

        // starts from vertex-0
        open_set.insert(0, 0);

        while (!open_set.empty()) {
            auto top = open_set.min_priority();
            auto top_idx = top.get_index();

            auto sub_arr = m_graph->sub_array(top_idx);
            for (int i = 0; i < vn; i++) {
                auto weight_top2i = (*sub_arr)[i];
                auto weight_sum_i = m_closed_set[i];
                // add adjacency vertex to open set
                if (m_closed_set[top_idx] > weight_sum_i + weight_top2i) {
                    m_closed_set[top_idx] = weight_sum_i + weight_top2i;
                    open_set.insert(i, m_closed_set[top_idx]);
                }

            }
        }
    }

    void Print() const {
        float weight_sum = .0f, valid_path_sum = .0f;
        for (auto &weight: m_closed_set) {
            if (weight) {
                weight_sum += static_cast<float>(weight);
                ++valid_path_sum;
            }
        }
        std::cout << "avg_weight = " << weight_sum << " / " << valid_path_sum << " = " <<
                  static_cast<float>(weight_sum / valid_path_sum) << std::endl;
    }

private:
    // graph
    Graph *m_graph;
    // closed set, contains vertices already expanded
    std::vector<uint> m_closed_set;
};

int main() {
    std::cout << "Hello, World!" << std::endl;

    std::vector<float> densities = {0.2f, 0.4f, 0.5f};
    for (auto &d: densities) {
        Graph graph(d, 50);
        auto dij = ShortestPath(graph);
        dij.Solve();
        dij.Print();
    }

    return 0;
}

