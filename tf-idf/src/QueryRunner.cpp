#include "../include/QueryRunner.hpp"
#include "../include/VectorMath.hpp"
#include <iostream>
#include <utility>
#include <algorithm>

QueryRunner::QueryRunner() {
    printf("Welcome to query runner :p \n");
}

std::vector<std::pair<std::string, float>> QueryRunner::run_query(std::vector<float> q ,std::pair<std::vector<std::string>, std::vector<std::vector<float>>> documents) {
    std::vector<std::pair<std::string, float>> result;

    for(int i = 0; i < documents.first.size(); i++) {
        std::vector<float> d = documents.second[i];

        float res = cosine_similarity(q, d);

        std::pair<std::string, float> curr_doc_relevancy = std::make_pair(documents.first[i], res);
        result.push_back(curr_doc_relevancy);
    }

    std::sort(result.begin(), result.end(), [](const auto &a, const auto &b) {
        return a.second > b.second;
    });

    return result;
}

float QueryRunner::cosine_similarity(std::vector<float> q, std::vector<float> d) {
    double dot_res = VectorMath::dot(q,d);

    double q_mag = VectorMath::magnitude(q);
    double d_mag = VectorMath::magnitude(d);

    if (q_mag == 0 || d_mag == 0) {
        return 0.0f;
    }

    float result = dot_res / (q_mag * d_mag);

    return result;
}
