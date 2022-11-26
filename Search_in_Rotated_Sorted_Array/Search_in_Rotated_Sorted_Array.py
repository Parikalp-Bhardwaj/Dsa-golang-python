nums = [4,5,6,7,0,1,2]
target = 0


def search(nums,target):
    start,end = 0,len(nums)-1

    while start <= end:
        mid = (start + end) //2

        if nums[mid] == target:
            return mid

        if nums[start] <= nums[mid]:
            if target < nums[start] or nums[mid] < target:
                start = mid + 1
            else:
                end = mid - 1

        else:
            if target > nums[end] or target < nums[mid]:
                end = mid - 1
            else:
                start = mid + 1

    
    return -1




print(search(nums,target))