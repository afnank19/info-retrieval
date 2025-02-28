#ifndef QUERY_RUNNER_H
#define QUERY_RUNNER_H

#include <vector>
#include <map>
#include <string>

class QueryRunner {
private:
    /* data */
    static float cosine_similarity(std::vector<float> q, std::vector<float> d);
public:
    QueryRunner(/* args */);
    static std::vector<std::pair<std::string, float>> run_query(std::vector<float> q ,std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents);
};

#endif