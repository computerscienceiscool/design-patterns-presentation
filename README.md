# Design Patterns in Go — Workshop (March 2025)

This repository is part of a hands-on workshop on **Design Patterns** using the Go programming language. It demonstrates canonical examples from each of the three main pattern categories:

- **Creational Patterns**
- **Structural Patterns**
- **Behavioral Patterns**

## 📁 Directory Structure

```bash
.
├── behavioral/
│   ├── observer/
│   └── strategy/
├── creational/
│   ├── builder/
│   ├── factory/
│   └── singleton/
├── structural/
│   ├── adapter/
│   ├── decorator/
│   └── facade/
├── slides/
│   ├── slides.md         # Markdown source for the presentation
│   ├── slides.html       # Rendered HTML
│   ├── main.go           # Slide server (Go)
│   └── ...               # Supporting files
├── main.go               # Optional CLI/demo entry point
├── go.mod                # Go module definition
```

## ▶️ How to Run Examples

Each pattern lives in its own folder. To run an individual example:

```bash
cd creational/factory
go run main.go
```

Or run the top-level `main.go` if you have a combined demo:

```bash
go run main.go
```

## 🖥️ Slides

You can view the slides in HTML form:

- Open [`slides/slides.html`](slides/slides.html) in a browser, **or**
- Run the Go-based slide server:
  
```bash
cd slides
go run main.go
```

## 🧠 What's Covered

### Creational
- **Singleton**: Ensure a class has only one instance.
- **Factory**: Create objects without specifying exact classes.
- **Builder**: Construct complex objects step-by-step.

### Structural
- **Adapter**: Match interfaces of different classes.
- **Decorator**: Add responsibilities dynamically.
- **Facade**: Provide a simplified interface to a complex system.

### Behavioral
- **Observer**: Notify dependents of state changes.
- **Strategy**: Define interchangeable algorithms.

## 🧰 Requirements

- Go 1.20 or higher
- Basic command-line familiarity

## 📜 License

MIT License. See `LICENSE` file (if included).

------

## 🎮 Bonus: Kahoot Quiz!

Test your knowledge of Go design patterns with our interactive 17-question [Kahoot quiz](https://create.kahoot.it/details/39baaffd-263e-4cd1-be29-e4507025571c)!

Perfect for workshops, study groups, or reviewing on your own.

➡️ [Click here to play the Kahoot](https://create.kahoot.it/details/39baaffd-263e-4cd1-be29-e4507025571c)


Enjoy learning Go and design patterns!
