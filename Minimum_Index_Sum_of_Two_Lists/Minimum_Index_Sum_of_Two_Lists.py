
# list1 = ["Shogun","Piatti","Tapioca Express","Burger King","KFC"]
# list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
list1 = ["happy","sad","good"]
list2 = ["sad","happy","good"]
def findRestaurant(list1,list2):
    l1 = {j:i for i,j in enumerate(list1)}
    l2 = {j:i for i,j in enumerate(list2)}
    d={}
    for i in l1:
        if i in l2:
      
            val = l1[i]+l2[i]
            d[i]=val

    m = min(list(d.values()))
    res = []
    for key,val in d.items():
        if val <= m:
            m=val
            res.append(key)
    return res

print(findRestaurant(list1,list2))