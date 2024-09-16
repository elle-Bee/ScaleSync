import tkinter as tk
from tkinter import messagebox
import requests

def fetch_users():
    response = requests.get("http://localhost:8080/users")
    if response.status_code == 200:
        messagebox.showinfo("Users", response.json())
    else:
        messagebox.showerror("Error", "Could not fetch users")

def create_user():
    user_data = {
        "name": name_entry.get(),
        "email": email_entry.get(),
        "password": password_entry.get()
    }
    response = requests.post("http://localhost:8080/users", json=user_data)
    if response.status_code == 200:
        messagebox.showinfo("Success", "User created successfully")
    else:
        messagebox.showerror("Error", "Failed to create user")

app = tk.Tk()
app.title("User Management")

tk.Label(app, text="Name:").pack()
name_entry = tk.Entry(app)
name_entry.pack()

tk.Label(app, text="Email:").pack()
email_entry = tk.Entry(app)
email_entry.pack()

tk.Label(app, text="Password:").pack()
password_entry = tk.Entry(app, show="*")
password_entry.pack()

tk.Button(app, text="Create User", command=create_user).pack()
tk.Button(app, text="Fetch Users", command=fetch_users).pack()

app.mainloop()
