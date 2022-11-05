import numpy as np 

lis = np.random.choice([4, 4+1], size=(2, 8, 2)).astype(dtype=np.int32)

print(lis)

print("0 0", lis[0,0,0])
print("0 0", lis[0,0,1])

print()

print("0 1", lis[0,1,0])
print("0 1", lis[0,1,1])

print()

print("1 1", lis[1,1,0])
print("1 1", lis[1,1,1])