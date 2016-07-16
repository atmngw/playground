# coding: utf-8
import time
import random

list = list(range(1,100))
random.shuffle(list)

length = len(list)
print (list)
flag = True

start = time.time()
while flag :
    flag = False
    for index in range(0,length) :
        if index >= 1 :
            if list[index] < list[index-1] :
                list[index],list[index-1] = list[index-1],list[index]
                flag = True

end = time.time()
print (list)
print ("sort: %f seconds" % (end - start))
