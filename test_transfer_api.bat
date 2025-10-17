@echo off
REM Test script for Transfer API
REM Make sure the server is running before executing this script

echo ============================================
echo Transfer API Test Script
echo ============================================
echo.

echo [1/7] Creating User 1 (Alice with 1000 points)...
curl -X POST http://localhost:3000/api/v1/users -H "Content-Type: application/json" -d "{\"name\":\"Alice\",\"email\":\"alice@example.com\",\"balance\":1000}"
echo.
echo.

echo [2/7] Creating User 2 (Bob with 0 points)...
curl -X POST http://localhost:3000/api/v1/users -H "Content-Type: application/json" -d "{\"name\":\"Bob\",\"email\":\"bob@example.com\",\"balance\":0}"
echo.
echo.

timeout /t 2 >nul

echo [3/7] Getting all users...
curl http://localhost:3000/api/v1/users
echo.
echo.

timeout /t 2 >nul

echo [4/7] Creating transfer: Alice sends 100 points to Bob...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":100,\"note\":\"Payment for service\",\"idempotency_key\":\"transfer-test-001\"}"
echo.
echo.

timeout /t 2 >nul

echo [5/7] Testing idempotency (sending same transfer again)...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":100,\"note\":\"Payment for service\",\"idempotency_key\":\"transfer-test-001\"}"
echo.
echo.

timeout /t 2 >nul

echo [6/7] Checking user balances...
echo Alice's balance:
curl http://localhost:3000/api/v1/users/1/balance
echo.
echo Bob's balance:
curl http://localhost:3000/api/v1/users/2/balance
echo.
echo.

timeout /t 2 >nul

echo [7/7] Viewing transaction history...
echo Alice's ledger:
curl http://localhost:3000/api/v1/users/1/ledger
echo.
echo.
echo Bob's ledger:
curl http://localhost:3000/api/v1/users/2/ledger
echo.
echo.

echo ============================================
echo Test completed!
echo ============================================
echo.
echo Additional test commands:
echo - Get all transfers: curl http://localhost:3000/api/v1/transfers
echo - Get specific transfer: curl http://localhost:3000/api/v1/transfers/transfer-test-001
echo - Filter by user: curl http://localhost:3000/api/v1/transfers?user_id=1
echo - Filter by status: curl http://localhost:3000/api/v1/transfers?status=completed
echo.

pause
