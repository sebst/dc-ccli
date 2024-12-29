import { LazyLog } from "@melloware/react-logviewer";

export default function Logs() {
  return (
    <div>
      <h1>Logs</h1>
      Lazylog:
      <LazyLog url="http://127.0.0.1:6867/logs" />
    </div>
  );
}
