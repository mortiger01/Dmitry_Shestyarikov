from fastapi import FastAPI
from fastapi.responses import JSONResponse
from app.utils import get_days_until_new_year

# Создаем экземпляр приложения FastAPI
app = FastAPI(
    title="New Year Counter API",
    description="API для подсчета дней до Нового года",
    version="1.0.0"
)

@app.get("/info")
async def get_info():
    """
    Эндпоинт для получения количества дней до Нового года.
    
    Returns:
        JSON объект с полем days_before_new_year
    """
    days = get_days_until_new_year()
    
    return JSONResponse(
        content={"days_before_new_year": days},
        status_code=200
    )

@app.get("/")
async def root():
    """
    Корневой эндпоинт для проверки работоспособности сервера.
    """
    return {"message": "New Year Counter API is running", "endpoints": ["/info"]}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)