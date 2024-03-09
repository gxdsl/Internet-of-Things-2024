# 校园直饮水系统

## 数据库

### 数据库结构

**WaterDispenser 表结构**

| **字段名**    | **类型** | **描述**   |
| ------------- | -------- | ---------- |
| `ID`          | uint     | 主键       |
| `Price`       | float64  | 价格 元/升 |
| `DispenserID` | string   | 饮水机ID   |
| `Model`       | string   | 型号       |
| `Location`    | string   | 安装位置   |

**User 表结构**

| **字段名** | **类型** | **描述** |
| ---------- | -------- | -------- |
| `ID`       | uint     | 主键     |
| `User`     | string   | 用户名   |
| `Password` | string   | 密码     |
| `Card`     | string   | 卡号     |
| `Balance`  | float64  | 余额     |

**Transaction 表结构**
| **字段名**        | **类型**  | **描述**                 |
| ----------------- | --------- | ------------------------ |
| `ID`              | uint      | 主键                     |
| `User`            | string    | 用户                     |
| `DispenserID`     | uint      | 饮水机ID                 |
| `Amount`          | float64   | 金额                     |
| `TransactionTime` | time.Time | 消费时间，默认为当前时间 |

**DispenserStatus 表结构**

| **字段名**    | **类型**  | **描述**                 |
| ------------- | --------- | ------------------------ |
| `ID`          | uint      | 主键                     |
| `DispenserID` | uint      | 饮水机ID                 |
| `Status`      | string    | 状态                     |
| `Temperature` | float64   | 水温                     |
| `TDS`         | float64   | TDS水质                  |
| `Flow`        | bool      | 是否出水状态             |
| `RecordTime`  | time.Time | 记录时间，默认为当前时间 |

