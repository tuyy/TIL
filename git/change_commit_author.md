### 이미 commit 한 작성자 수정
```
git rebase -i HEAD~2  // 해당 커밋 'edit'
git commit --amend --author="사용자명 <이메일>"
git rebase --continue
```
