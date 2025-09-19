Write-Host "Starting DevToolBox Development Environment..." -ForegroundColor Green

Write-Host "Building Docker images..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml build

Write-Host "Starting services..." -ForegroundColor Yellow
docker-compose -f docker-compose.dev.yml up -d

Write-Host "Waiting for services to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 10

Write-Host "DevToolBox is running!" -ForegroundColor Green
Write-Host ""
Write-Host "Frontend: http://localhost:3000" -ForegroundColor Cyan
Write-Host "Backend API: http://localhost:8080" -ForegroundColor Cyan
Write-Host "PostgreSQL: localhost:5433" -ForegroundColor Cyan
Write-Host ""
Write-Host "Available commands:" -ForegroundColor Yellow
Write-Host "  make docker-logs    - View logs" -ForegroundColor White
Write-Host "  make docker-dev-down - Stop services" -ForegroundColor White
Write-Host "  make docker-clean   - Clean up everything" -ForegroundColor White
Write-Host ""
Write-Host "Happy coding!" -ForegroundColor Green
