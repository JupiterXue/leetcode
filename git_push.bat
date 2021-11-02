@echo off
set user=JupiterXue
set token=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
git add .
git commit -m "commit passed code"
echo %user%
echo %token%
git push
pause