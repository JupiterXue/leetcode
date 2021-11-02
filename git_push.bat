@echo off
set user=JupiterXue
set token=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
git add .
git commit -m "commit passed code"
echo %user%
echo %token%
git push
pause