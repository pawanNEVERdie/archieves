import React, { useState, useEffect } from "react";
import "./App.css";
export function App() {
  const [blogs, setBlogs] = useState([]);
  useEffect(() => {
    const getBlogs = async () => {
      try {
        const response = await fetch("http://localhost:3001/api/blogs");
        if (!response.ok) {
          throw new Error("Error loading blogs");
        }
        const blogs = await response.json();
        setBlogs(blogs);
      } catch (error) {
        console.log(error.message);
      }
    };
    getBlogs();
  }, []);
  return (
    <div className="app">
      <header className="header">
        <h2 className="title">Goxxygen Blog</h2>
      </header>
      <main className="main-content">
        {blogs &&
          blogs.map((blog) => (
            <div className="blog-card" key={blog.id}>
              <div className="blog-cover">
                <img src={blog.coverURL} alt={blog.title} />
              </div>
              <h3 className="blog-title">{blog.title}</h3>
            </div>
          ))}
      </main>
    </div>
  );
}