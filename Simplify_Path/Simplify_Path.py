# path = "/home/"
path = "/home//foo/"

# def simplifyPath(path):
#     path = path.split("/")
#     ls = []
#     for pa in path:
#         if pa == "." or pa == "":
#             continue
#         if pa == "..":
#             if ls:
#                 ls.pop()
#         else:
#             ls.append(pa)

#     return "/"+"/".join(ls)


# print(simplifyPath(path))


def simplifyPath(path):
    def rec(path,s1,p1):
        # if path == "":
        #     return s1
        print(len(path),len(s1))
        # for p in p1:
        #     if p == "." or p == "":
        #         continue
        #     if p =="..":
        #         if s1:
        #             s1.pop()
        #     else:
        #         s1.append(p)

        # print(s1)

        

    return rec(path,[],path.split("/"))


print(simplifyPath(path))