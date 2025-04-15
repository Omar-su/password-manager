import React, { useState } from "react";
import "./App.css";

function App() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
    site: "",
  });

  const handleChange = (e) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/api/pw", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    });

    const data = await response.json();
    alert(data.status);
  };

  return (
    <div className="container">
      <h2>Password Saver</h2>
      <form onSubmit={handleSubmit}>
        <input
          name="username"
          placeholder="Username"
          value={formData.username}
          onChange={handleChange}
          required
        />
        <input
          name="password"
          type="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange}
          required
        />
        <input
          name="site"
          placeholder="Site or App"
          value={formData.site}
          onChange={handleChange}
          required
        />
        <button type="submit">Save</button>
      </form>
    </div>
  );
}

export default App;
