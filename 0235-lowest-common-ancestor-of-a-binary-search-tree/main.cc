/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        TreeNode * cur = root;
        if (cur->val >= p->val && cur->val <=q->val) {
            return cur;
        }
        if (cur->val >=q->val && cur->val <=p->val) {
            return cur;
        }
        if (cur->val > p->val && cur->val > q->val) {
            return lowestCommonAncestor(cur->left, p, q);
        }
        if (cur->val < p->val && cur->val < q->val) {
            return lowestCommonAncestor(cur->right, p, q);
        }
    }
};
