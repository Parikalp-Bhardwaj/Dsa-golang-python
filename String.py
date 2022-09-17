# s = "book"
# s = "textbook"
# s="AbCdEfGh"

# s = "Uo"
# s="cVQMDfuraWDVrMUKICHXyCqyDzbcYXBHgVxCIlVbhWtYRvvpYDPzLkqIpBRPBRZjkWkPAkfgkNTMBJnpzrFuJvDsycbxBzswEMWxJisWgQmzYLhPBttjdofMhsGJTbqChdgbEdyEYhqUHCdJPSLG"
# def halvesAreAlike(s):
#     vowels = "aeiouAEIOU"
#     spl = len(s) // 2
#     a = s[:spl]
#     b = s[spl:]

#     i=0
#     a1 = 0
#     b1 = 0
#     while len(a) > i and len(b) > i:     
#         if a[i] in vowels:
#             a1 +=1
#         if b[i] in vowels:
#             b1 +=1
#         i+=1

#     if a1 == b1:
#         return True
#     return False

# print(halvesAreAlike(s))




# s = "a1c1e1"

# def replaceDigits(s):
#     ls = "".join([chr(ord(s[i-1]) + int(s[i])) if s[i].isdigit() else s[i] for i in range(len(s)) ])
#     print(ls)

# print(replaceDigits(s))


# nums = [1,2,3,4]

# # nums = [1,1,2,3]
# def decompressRLElist(nums):
#     # nums = [nums[i] for i in range(len(nums)) if i%2]
#     # print(nums)

#     ls = []
#     for i in range(0,len(nums),2):
#         num = nums[i]
#         val = [nums[i+1]]*num
#         ls.append(val)
#     return sum(ls,[])

#     # count = 1 
#     # lst = []
#     # for i in nums[::2]:
#     #     l = [nums[count]]*i
#     #     print(l)
#     # #     lst.append(l)
#     #     count += 2
#     # print(lst)

# print(decompressRLElist(nums))


# nums = [1,2,2,1]
# k = 1

# def countKDifference(nums,k):
#     # nums.sort()
#     count = 0
#     for i in range(len(nums)):
#         for j in range(i+1,len(nums)):
#             if abs(nums[j] - nums[i]) == k:
#                 count+=1

#     # print(count)
# print(countKDifference(nums,k))




# nums = [5,6,2,7,4]
# def maxProductDifference(nums):
#     nums.sort()
#     return (nums[-1] * nums[-2] - nums[0] * nums[1] )

#     # if nums[0] > nums[1]:
#     #     smallet = nums[1]
#     #     secondSmallest = nums[0]
#     # else:
#     #     smallet,secondSmallest = nums[0],nums[1]
    
#     # for i in range(2,len(nums)):
#     #     if nums[i] < smallet:
#     #         secondSmallest = smallet
#     #         smallet = nums[i]
#     #     elif nums[i] >= smallet and nums[i] < secondSmallest:
#     #         secondSmallest = nums[i]

#     # print(smallet,secondSmallest)


# print(maxProductDifference(nums))



# nums = [1,1,1]
# nums = [1,5,2,4,1]

# def minOperations(nums):
    
#     count = 0
    # for i in range(1,len(nums)):
    #     if nums[i] <= nums[i-1]:
    #         count += (nums[i-1] - nums[i])+1
    #         nums[i] += (nums[i-1] - nums[i])+1

    # print(count)

#     i=1
#     while len(nums) > i:
#         if nums[i] <= nums[i-1]:
#             count += (nums[i-1] - nums[i])+1
#             nums[i] += (nums[i-1] - nums[i])+1
#         i+=1

#     return count


  
# print(minOperations(nums))

# arr = [17,18,5,4,6,1]
# arr = [400]
# def replaceElements(arr):
#     len(arr)
#     print(len(arr))

#     ls= []
#     if len(arr) > 1:
#         for i in range(1,len(arr)-1):
#             max_val = max(arr[i:])
#             ls.append(max_val)

#         ls.append(arr[-1])
#     ls.append(-1)
#     print(ls)
# print(replaceElements(arr))

# startTime = [1,2,3]
# endTime = [3,2,7]
# queryTime = 4


# def busyStudent(startTime,endTime,queryTime):
#     count = 0
#     for i,j in zip(startTime,endTime):
#         if i <= queryTime and j >= queryTime:
#             count +=1

#     print(count)

# print(busyStudent(startTime,endTime,queryTime))




# boxTypes = [[1,3],[2,2],[3,1]]
# truckSize = 4

# boxTypes = [[5,10],[2,5],[4,7],[3,9]]
# truckSize = 10


# def maximumUnits(boxTypes,truckSize):
    
#     boxTypes.sort(key=lambda x:x[1], reverse = True)
#     d = {j:i for i,j in boxTypes}
#     print(d)
#     c = 0
#     rem = 0
#     mul = 0
#     # for i,(key,val) in enumerate(d.items()):
#     #     c+=val
#     #     mul += key*val
#     #     print("mul out",key,val,mul)
#     #     if c >= truckSize:
#     #         print("inner ",key,val,mul)
#     #         mul -= key*val
#     #         print("inner red",key,val,mul)
#     #         rem = c - truckSize
#     #         c -= rem
#     #         if c == truckSize:
#     #             print("c ",c)
#     #             print("ok")
#     #             mul+= val*rem
#     #             print("ke ",mul)
#     #             break

#     print(boxTypes)
#     count = 0
#     total = 0
#     for box in boxTypes:
#         total += box[0]
#         count += box[0]*box[1]
#         if total >= truckSize:
#             count -= box[0]*box[1]
#             rem = total - truckSize
            
#             count += box[1]*rem
#             mul += box[1] * rem
#             print("rem ",rem)
#             print("total ",total)
#             print("count ",count)
#             break
#         # print("total ",total)
            
#     print(mul)

# print(maximumUnits(boxTypes,truckSize))






