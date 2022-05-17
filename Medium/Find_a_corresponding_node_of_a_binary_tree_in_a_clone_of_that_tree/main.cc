// https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/
#include <bits/stdc++.h>

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution
{
public:
  TreeNode *getTargetCopy(TreeNode *original, TreeNode *cloned,
                          TreeNode *target)
  {
    if (original == target || original == NULL)
      return cloned;
    TreeNode *left  = getTargetCopy(original->left, cloned->left, target),
             *right = getTargetCopy(original->right, cloned->right, target);
    return left ? left : right;
  }
};
