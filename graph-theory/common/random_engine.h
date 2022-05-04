//
// Created by gorjess on 5/4/2022.
//

#ifndef GRAPH_THEORY_COMMON_RANDOM_ENGINE_H
#define GRAPH_THEORY_COMMON_RANDOM_ENGINE_H

#include <random>

template<typename T>
class RandomEngine
{
public:
    RandomEngine(T a, T b): m_a(a), m_b(b)
    {
        init();
    }

    explicit RandomEngine(T upper_bound): m_a(0), m_b(upper_bound)
    {
        init();
    }

    ~RandomEngine() = default;

    // Generate a random number between m_a and m_b.
    int Gen() { return m_dist(m_gen); }

private:
    inline void init()
    {
        std::random_device dev;
        m_gen = std::mt19937_64(dev());
        m_dist = std::uniform_int_distribution<T>(m_a, m_b);
    }

private:
    T m_a, m_b;
    std::mt19937_64 m_gen;
    std::uniform_int_distribution<int> m_dist;
};


#endif //GRAPH_THEORY_COMMON_RANDOM_ENGINE_H
