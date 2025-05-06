CXX = g++
CXXFLAGS = -Iinclude -Iporter2_stemmer
SRC = ./src/*.cpp porter2_stemmer/porter2_stemmer.cpp
OUT = main

run: $(SRC)
	$(CXX) $(CXXFLAGS) $(SRC) -o $(OUT)
	./$(OUT)

build: $(SRC)
	$(CXX) $(CXXFLAGS) $(SRC) -o $(OUT)

run-build:
	./$(OUT)

clean:
	rm ./$(OUT)