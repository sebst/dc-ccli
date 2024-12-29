// import { LazyLog } from "@melloware/react-logviewer";

import StreamingLogsComponent from "@/components/streaming-logs-component";

export default function Logs() {
  return (
    <div>
      <h1>Logs</h1>
      Lazylog:
      {/* <LazyLog url="http://127.0.0.1:6867/logs" /> */}
      <StreamingLogsComponent />
    </div>
  );
}
