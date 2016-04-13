#ifndef SYNACORE_MEMORY_H
#define SYNACORE_MEMORY_H

#include <iostream>
#include <stdint.h>

class memory {
    uint32_t memoryLength = 6;
    uint32_t *memory;
    uint32_t memoryPointer = 0;
public:
    void loadFromFile(std::string fileName);

    uint32_t getCurrentMemory();

    uint32_t getPointer();
    void setPointer(uint32_t pointer);

    bool isEOM();
};


#endif //SYNACORE_MEMORY_H
