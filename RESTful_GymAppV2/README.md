Here’s a **complete step-by-step guide** for anyone to take your uploaded project folder and run it fully, including database setup and testing with the frontend:

---

## 1. Download the Project

1. Download the project ZIP from GitHub or any source.
2. Extract the folder. Assume the folder is named `gym-management`.

---

## 2. Install Requirements

* **Go** (1.18+ recommended)
  Check installation:

```bash
go version
```

* **PostgreSQL**
  Make sure it’s running and you can create databases.

---

## 3. Configure Database

1. Open `database.env` and set your PostgreSQL credentials:

```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gym
```

2. Create the database `gym` (or whatever name you put in `DB_NAME`):

```sql
CREATE DATABASE gym;
```

3. Create tables using PostgreSQL (via psql or pgAdmin):

```sql
-- Users table
CREATE TABLE user_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    lastname VARCHAR(100),
    age INT,
    height INT,
    weight FLOAT,
    gender VARCHAR(10),
    skilllevel VARCHAR(20),
    plan VARCHAR(30),
    email VARCHAR(100)
);

-- Coaches table
CREATE TABLE coach_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    lastname VARCHAR(100),
    age INT,
    gender VARCHAR(10),
    email VARCHAR(100),
    cost FLOAT,
    verified BOOLEAN DEFAULT FALSE,
    profile_pic VARCHAR(255),
    bio VARCHAR(5000)
);
```

---

## 4. Install Go Dependencies

Open a terminal in the project folder:

```bash
go mod tidy
```

This will download all required packages (`gin`, `uuid`, `pq`, etc.).

---

## 5. Prepare Uploads Folder

Ensure the folder for coach profile pictures exists:

```bash
mkdir -p uploads/coaches
```

---

## 6. Run the Backend

```bash
go run main.go
```

* By default, the API will run on `http://localhost:8080`.

---

## 7. Open the Frontend

1. Navigate to the `html` folder.
2. Open `index.html` in a browser.
3. Use the **role selector** to switch between `Trainee` and `Coach`.
4. Fill in forms to **add/edit/delete users and coaches**.
5. Upload profile pictures for coaches if needed.

All actions will update the database in real-time.

---

## 8. Notes / Testing Tips

* Profile pictures are stored in `uploads/coaches/` with coach ID filenames.
* Coach biography field (`bio`) can store up to 5000 characters.
* All IDs are UUIDs, generated automatically.
* To reset the database, drop and recreate tables.
* Make sure `database.env` matches your PostgreSQL setup.

