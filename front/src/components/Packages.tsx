import { useState } from "react";

const url = "http://127.0.0.1:6867/api/test";

export default function PackagesComponent() {
  const [logLines, setLogLines] = useState<string[]>([]);

  async function clickAdd() {
    try {
      const response = await fetch(url);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const reader = response.body?.getReader();
      const decoder = new TextDecoder("utf-8");
      let partialLine = "";

      while (reader) {
        const { value, done } = await reader.read();
        if (done) {
          break; // If the stream is done, exit the loop
        }
        const chunk = decoder.decode(value, { stream: true });
        partialLine += chunk;

        const lines = partialLine.split("\n");
        partialLine = lines.pop() || ""; // Keep the incomplete line

        for (const line of lines) {
          setLogLines((prev) => [...prev, line]);
        }
      }
      if (partialLine.length > 0) {
        setLogLines((prev) => [...prev, partialLine]);
      }
      setLogLines((prev) => [...prev, "END"]);
    } catch (error) {
      console.error("Error fetching and processing the stream:", error);
    }
  }

  return (
    <>
      Packages
      {/* Two divs, side by side, with tailwind */}
      <div className="flex">
        <div className="w-1/3 ">
          <input type="text" className="w-full" placeholder="package" />
        </div>
        <div className="w-1/3 ">
          <input type="text" className="w-full" placeholder="version" />
        </div>
        <div className="w-1/3 ">
          <button className="w-full   " onClick={clickAdd}>
            Add
          </button>
        </div>
      </div>
      <code>
        <pre>
          {logLines.map((line, index) => (
            <div key={index}>{line}</div>
          ))}
        </pre>
      </code>
    </>
  );
}
