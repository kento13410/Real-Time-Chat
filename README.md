# Real-Time-Chat

このプロジェクトは、2月の春休み期間を利用して、Go言語、Redis、およびgRPCを用いてリアルタイムでの通信が可能なチャットアプリケーションを開発することを目的としています。

## 技術スタック

- **言語**: Go
- **データストア**: PostgresSQL, Redis
- **ORM**: ent
- **通信**: gRPC

## 主な機能

- ユーザー間でのリアルタイムメッセージング

## システム設計
### フローチャート
```mermaid
flowchart TD
    A[Start] --> B(Signup)
    B --> C(Login)
    C --> D{Home Screen}
    D --> E[Search by Username]
    E --> F[Add Friend]
    F --> D
    D --> G[Tap on Friend]
    G --> H[Chat Screen]
    H --> I[Display Past Chats]
    I --> D
```

### データベース設計
```mermaid
erDiagram
    USERS ||--o{ USER_RELATIONS : has
    USERS {
        int id PK "Primary Key"
        string username "Unique"
        string email "Unique"
        string password_hash
    }
    USER_RELATIONS ||--o{ CHAT : communicates
    USER_RELATIONS {
        int id PK "Primary Key"
        int user_id_1 FK "Foreign Key to USERS.id"
        int user_id_2 FK "Foreign Key to USERS.id"
    }
    CHAT {
        int id PK "Primary Key"
        int sender_id FK "Foreign Key to USERS.id"
        int receiver_id FK "Foreign Key to USERS.id"
        string message
        datetime sent_at
    }
```