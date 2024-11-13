import { useEffect, useRef, useState } from "react";
import { baseURL } from "../common/urls";

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
  }, [value, delay]); // Now runs every time value or delay changes

  return delayedValue;
}

function useDebounce<T>(value: T, delay: number) {
  const [delayedValue, setDelayedValue] = useState(value);

  useEffect(() => {
    const timeoutID = setTimeout(() => {
      setDelayedValue(value);
    }, delay);

    return () => clearTimeout(timeoutID);
  }, [value, delay]); // Now runs every time value or delay changes

  return delayedValue;
}

export function SynchEditor() {
  const [clientContent, setClientContent] = useState("");
  const [stageContent, setStageContent] = useState("");
  const [serverContent, setServerContent] = useState("");
  const delayedValue = useDelay(clientContent, 500);
  const [isReadyToSynchronizeChanges, setIsReadyToSynchronizeChanges] = useState(false)
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
      if (data.ack) {
        console.log("Recieved acknowledgement");
        setServerContent(stageContent);
        return;
      }
      if (data.content) {
        console.log("Recieved base content")
        const content = linesToString(data.content);
        setServerContent(content);
        setIsReadyToSynchronizeChanges(true)
        return;
      }
      if (data.diff) {
        console.log(data.diff)
        const newLines = linesToString(
          addDifference(stringToLines(serverContent), data.diff)
        );
        console.log(newLines)
        setServerContent(newLines);
        setIsReadyToSynchronizeChanges(true);
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
    setClientContent(e.target.value);
  }

  useEffect(() => {
    if (ws.current) {
      const clientLines = stringToLines(clientContent);
      const serverLines = stringToLines(serverContent);
      const data = {
        diff: compareLines(clientLines, serverLines).map((index) => {
          return {
            index: index,
            data: stringToLines(clientContent)[index],
          };
        }),
      };
      if (data.diff.length == 0) return;
      setStageContent(clientContent);
      ws.current.send(JSON.stringify(data));
    }
  }, [delayedValue]);

  useEffect(() => {
    if (isReadyToSynchronizeChanges) {
        setClientContent(serverContent)
        setIsReadyToSynchronizeChanges(false)
    }
  }, [isReadyToSynchronizeChanges]);

  return (
    <>
      <div>
        <textarea
          value={clientContent}
          onChange={handleChange}
          style={{ width: "300px", height: "200px" }}
        ></textarea>
      </div>
      <div>
        <button onClick={connectToWS}>connect</button>
      </div>
      <div>
        <button onClick={disconnectFromWS}>disconnect</button>
      </div>
    </>
  );
}

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

interface Difference {
  index: number;
  data: string;
}

function addDifference(lines: string[], diff: Difference[]) {
  for (let i = 0; i < diff.length; i++) {
    if (diff[i].index >= lines.length) {
      lines.push(diff[i].data);
    } else {
      lines[diff[i].index] = diff[i].data;
    }
  }
  return lines;
}
