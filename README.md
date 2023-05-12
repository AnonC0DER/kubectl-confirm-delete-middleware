# kubectl delete confirmation script

It's a really simple script with less than 50 lines of code yet, in my opinion, 
a useful middleware to prevent irreversible mistakes.

### How to use?
- `go build`
- 'mv kubectl-confirm-delete /usr/share/mydir/'
- Add this alias to your zshrc or bashrc
- `alias kubectl="/usr/share/mydir/kubectl-confirm-delete"`

Then, if you run "kubectl delete whatever_resource whatever_name", you should be able to see 
the confirmation message.