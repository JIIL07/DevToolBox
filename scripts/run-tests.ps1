Write-Host "Running all tests for DevToolBox..." -ForegroundColor Green

Write-Host ""
Write-Host "Testing Go backend..." -ForegroundColor Yellow
go test ./... -v

Write-Host ""
Write-Host "Testing TypeScript frontend..." -ForegroundColor Yellow
Set-Location frontend
npm test -- --run
Set-Location ..

Write-Host ""
Write-Host "Testing Python plugins..." -ForegroundColor Yellow
Set-Location tests/python
python -m pytest -v
Set-Location ../..

Write-Host ""
Write-Host "All tests completed successfully!" -ForegroundColor Green