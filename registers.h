#ifndef SYNACORE_REGISTERS_H
#define SYNACORE_REGISTERS_H

#include <stdint.h>

class registers {
    uint32_t registers[8] = {0, 0, 0, 0, 0, 0, 0, 0};
public:
    uint32_t getRegister(uint32_t index);
    void setRegsiter(uint32_t index, uint32_t value);
};


#endif //SYNACORE_REGISTERS_H
