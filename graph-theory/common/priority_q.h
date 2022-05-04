#include <queue>
#include <unordered_map>


template<typename entry_type>
class PriorityQ {
public:
    // changes the priority of queue element
    void change_priority(entry_type& entry) {
        auto iter = m_vertices.find(idx);
        if (iter != m_vertices.end())
            if (iter->second->get_value() <= value)
                return;
        iter->second->set_expired();

        m_pq.push(entry);  // todo: try use emplace.
        m_vertices.insert(entry);
    }

    // removes and returns the top element of the queue.
    entry_type min_priority(bool remove = true) {
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
    void insert(entry_type& entry) {
        change_priority(entry);
    }

    // returns the top element of the queue.
    entry_type top() {
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
    std::unordered_map<uint, entry_type *> m_vertices;
};
