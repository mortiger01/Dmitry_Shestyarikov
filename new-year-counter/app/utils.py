from datetime import datetime, timezone
from dateutil import tz

def get_days_until_new_year():
    """
    Рассчитывает количество дней до следующего Нового года.
    Учитывает часовой пояс UTC для корректного расчета.
    """
    # Получаем текущую дату и время в UTC
    now = datetime.now(timezone.utc)
    
    # Определяем следующий Новый год (1 января следующего года)
    next_new_year = datetime(now.year + 1, 1, 1, tzinfo=timezone.utc)
    
    # Рассчитываем разницу в днях
    days_remaining = (next_new_year - now).days
    
    return days_remaining