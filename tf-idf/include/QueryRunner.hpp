#ifndef QUERY_RUNNER_H
#define QUERY_RUNNER_H

#include <vector>
#include <map>
#include <string>
#include "Indexer.hpp"

class QueryRunner {
private:
    // Indexer indexer;
    std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents;
    /* data */
    static float cosine_similarity(std::vector<float> q, std::vector<float> d);
public:
    QueryRunner();
    static std::vector<std::pair<std::string, float>> run_query(std::vector<float> q ,std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents);
    std::vector<std::pair<std::string, float>> run_query_v2(std::vector<float> q);
};

#endif