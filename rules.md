func AppendToFile(filename string, text string) error {
    // 1. Open the file:
    // O_WRONLY: Open for writing only
    // O_CREATE: Create it if it doesn't exist
    // O_APPEND: Start writing at the end of the file
    f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    
    // 2. Always defer closing the file so it cleans up after itself
    defer f.Close()

    // 3. Write the text (plus a newline so the next write starts on a fresh line)
    _, err = f.WriteString(text + "\n")
    return err
}