import { useEffect, useRef, useState } from "react";
import { baseURL } from "../common/urls";

function linesToString(lines: string[]): string {
  return lines.join("\n");
}

function stringToLines(str: string): string[] {
  return str.split("\n");
}

function compareLines(lines1: string[], lines2: string[]): number[] {
  const diff: number[] = [];
  const smallerLength = Math.min(lines1.length, lines2.length);
  const biggerLength = Math.max(lines1.length, lines2.length);

  for (let i = 0; i < biggerLength; i++) {
    if (i >= smallerLength || lines1[i] != lines2[i]) {
      diff.push(i);
    }
  }
  return diff;
}

function cursorPositionToLineIndex(pos: number, lines: string[]): number {
  let count = 0;
  for (let i = 0; i < lines.length; i++) {
    const len = lines[i] == "" ? 1 : lines[i].length + 1;
    if (pos <= count + len) {
      return i;
    }
    count += len;
  }
  return -1;
}

interface Difference {
  index: number;
  data: string;
}

function addDifference(lines: string[], diff: Difference[]) {
  console.log(lines);
  for (let i = 0; i < diff.length; i++) {
    while (diff[i].index >= lines.length) {
      lines.push("");
    }
    lines[diff[i].index] = diff[i].data;
  }
  console.log(lines);
  return lines;
}

function useDelay<T>(value: T, delay: number) {
  const [delayedValue, setDelayedValue] = useState(value);
  const timeoutID = useRef(0);

  useEffect(() => {
    if (timeoutID.current) return;
    timeoutID.current = setTimeout(() => {
      setDelayedValue(value);
      clearTimeout(timeoutID.current);
      timeoutID.current = 0;
    }, delay);
  }, [value, delay]);

  return delayedValue;
}
export function SynchEditor() {
  const [currentCursorPosition, setCurrentCursorPosition] = useState(0);
  const [lockedLines, setLockedLine] = useState<Map<number, string>>(
    new Map<number, string>()
  );
  const [clientContent, setClientContent] = useState("");
  const [serverContent, setServerContent] = useState("");
  const delayedValue = useDelay(clientContent, 250);
  const [isBusy, setIsBusy] = useState(false);
  const [shouldSyncronize, setShouldSyncronize] = useState(true);
  const ws = useRef<WebSocket>();

  function connectToWS() {
    ws.current = new WebSocket(baseURL + "/editor/open");

    ws.current.addEventListener("open", () => {
      console.log("WebSocket connection opened");

      ws.current!.send(
        JSON.stringify({
          token: "b9d30a2c-023b-4587-bf62-fde58fa7baa6",
        })
      );
      ws.current!.send(
        JSON.stringify({
          filename: "test.py",
        })
      );
    });

    ws.current.addEventListener("close", (e) => {
      console.log("WebSocket closed", e.target);
    });

    ws.current.addEventListener("message", (e) => {
      const data = JSON.parse(e.data);
      if (data.error_message) {
        console.error(data);
        return;
      }
      if (data.message) {
        console.log(data.message);
        return;
      }
      if (data.locked) {
        setLockedLine((old) => {
          const newMap = new Map(old);
          newMap.delete(data.unlocked - 1);
          newMap.set(data.locked - 1, data.by);
          return newMap;
        });
        return;
      }
      if (data.content) {
        setServerContent(data.content);
        setShouldSyncronize(true);
        return;
      }
      console.log("Recieved unknown data", data);
    });
  }

  function disconnectFromWS() {
    if (!ws.current) return;
    ws.current.close(1000);
    ws.current = undefined;
  }

  function handleChange(e: React.ChangeEvent<HTMLTextAreaElement>) {
    const currentCursorPosition = cursorPositionToLineIndex(
      e.target.selectionStart,
      stringToLines(e.target.value)
    );
    setCurrentCursorPosition(currentCursorPosition);

    setIsBusy(true);
    setClientContent(e.target.value);
  }

  useEffect(() => {
    if (ws.current) {
      ws.current.send(
        JSON.stringify({
          content: clientContent,
        })
      );
    }
  }, [delayedValue]);

  useEffect(() => {
    if (ws.current) {
      ws.current.send(
        JSON.stringify({
          cursor_position: currentCursorPosition + 1,
        })
      );
    }
  }, [currentCursorPosition]);

  useEffect(() => {
    console.log("Cursor position changed", currentCursorPosition);
  }, [currentCursorPosition]);

  useEffect(() => {
    const timeoutID = setTimeout(() => {
      setIsBusy(false);
    }, 500);
    return () => clearTimeout(timeoutID);
  }, [clientContent]);

  useEffect(() => {
    if (!isBusy && shouldSyncronize) {
      setClientContent(serverContent);
      setShouldSyncronize(false);
    }
  }, [isBusy, shouldSyncronize]);

  return (
    <>
      <div>
        <textarea
          value={clientContent}
          onChange={(e) => {
            handleChange(e);
          }}
          style={{ width: "300px", height: "200px", lineHeight: "15px" }}
        ></textarea>
      </div>
      {Array.from(lockedLines.entries()).map(([key, value]) => (
        <div
          key={key}
          style={{ position: "absolute", left: "310px", top: 100 + key * 15 }}
        >
          &lt;== {value}
        </div>
      ))}
      <div>
        <button onClick={connectToWS}>connect</button>
      </div>
      <div>
        <button onClick={disconnectFromWS}>disconnect</button>
      </div>
    </>
  );
}
