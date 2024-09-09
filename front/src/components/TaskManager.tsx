/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useState, useEffect } from 'react';
import { Sparklines, SparklinesLine } from 'react-sparklines';

const url = 'http://127.0.0.1:8080/api/processes'

const TaskManager: React.FC = () => {
  interface Process {
    PID: number;
    Name: string;
    CPUPercent: number;
    MemoryUsage: number;
    User: string;
    StartTime: string;
  }
  
  const [processes, setProcesses] = useState<Process[] | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [memoryHistory, setMemoryHistory] = useState<{ [key: number]: number[] }>({});
  const [cpuHistory, setCpuHistory] = useState<{ [key: number]: number[] }>({});

  const cpuHist = (data: any[]) => {
    data.forEach((process) => {
        if (cpuHistory[process.PID]) {
            cpuHistory[process.PID].push(process.CPUPercent);
        } else {
            cpuHistory[process.PID] = [process.CPUPercent];
        }
    });
    setCpuHistory(cpuHistory);
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const memHist = (data: any[]) => {
    data.forEach((process) => {
        if (memoryHistory[process.PID]) {
            memoryHistory[process.PID].push(process.MemoryUsage);
        } else {
            memoryHistory[process.PID] = [process.MemoryUsage];
        }
    });
    setMemoryHistory(memoryHistory);
}

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error(`Error: ${response.statusText}`);
        }
        const jsonData = await response.json();
        memHist(jsonData);
        cpuHist(jsonData);
        setProcesses(jsonData);
        setError(null);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message);
        }
        setProcesses(null);
      }
    };

    // Fetch the data immediately, then every 2 seconds
    fetchData();
    const intervalId = setInterval(fetchData, 2000);

    // Cleanup interval on component unmount
    return () => clearInterval(intervalId);
  }, []);

  if (error || !processes) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
            <h1>Task Manager</h1>


            <table>
                <thead>
                    <tr>
                        <th>PID</th>
                        <th>Command</th>
                        <th>CPUPercent</th>
                        <th>MemoryUsage</th>
                        <th>User</th>
                        <th>StartTime</th>
                    </tr>
                </thead>
                <tbody>
                    {processes.map((process) => {
                        return (
                            <tr key={process.PID}>
                                <td>{process.PID}</td>
                                <td>{process.Name}</td>
                                <td>
                                    {/* {process.CPUPercent} */}
                                    {/* <br /> */}
                                    <Sparklines data={cpuHistory[process.PID]} limit={5} width={100} height={20} margin={5}>
                                        <SparklinesLine color="blue" />
                                    </Sparklines>
                                </td>
                                <td>
                                    {/* {process.MemoryUsage} */}
                                    {/* <br /> */}
                                    <Sparklines data={memoryHistory[process.PID]} limit={5} width={100} height={20} margin={5}>
                                        <SparklinesLine color="blue" />
                                    </Sparklines>
                                </td>
                                <td>{process.User}</td>
                                <td>{process.StartTime}</td>
                            </tr>
                        );
                    })}
                </tbody>
            </table>
        </div>
  );
};

export default TaskManager;
