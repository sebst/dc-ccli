// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { useEffect, useState } from "react";
import { Sparklines, SparklinesLine } from 'react-sparklines';

const url = "http://127.0.0.1:3002/api/processes"


async function getProcesses() {
    console.log("Fetching processes");
    const response = await fetch(url);
    const data = await response.json();
    return data;
}

function TaskManager() {

    const [processes, setProcesses] = useState([]);
    const [memoryHistory, setMemoryHistory] = useState<{ [key: number]: number[] }>({});
    const [cpuHistory, setCpuHistory] = useState<{ [key: number]: number[] }>({});

    interface ProcessData {
        PID: number;
        CPUPercent: number;
        Name: string;
        MemoryUsage: number;
        User: string;
        StartTime: string;
    }
    
    const cpuHist = (data: ProcessData[]) => {
        data.forEach((process: ProcessData) => {
            if (cpuHistory[process.PID]) {
                cpuHistory[process.PID].push(process.CPUPercent);
            } else {
                cpuHistory[process.PID] = [process.CPUPercent];
            }
        });
        setCpuHistory(cpuHistory);
    }

    const memHist = (data: ProcessData[]) => {
        data.forEach((process: ProcessData) => {
            if (memoryHistory[process.PID]) {
                memoryHistory[process.PID].push(process.MemoryUsage);
            } else {
                memoryHistory[process.PID] = [process.MemoryUsage];
            }
        });
        setMemoryHistory(memoryHistory);
    }

    const updateProcesses = () => {
        getProcesses().then(data => {
            setProcesses(data);
            cpuHist(data);
            memHist(data);
        });
    }

    // const updateProcessesRegularly = () => {
    //     setInterval(() => {
    //         updateProcesses();
    //     }, 500);
    // }

    // useEffect(() => {
    //     updateProcessesRegularly();
    // }, [processes]);


    return (
        <div>
            <h1>Task Manager</h1>

            <button onClick={updateProcesses}>Update</button>
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
                    {processes.map((process: ProcessData) => {
                        return (
                            <tr key={process.PID}>
                                <td>{process.PID}</td>
                                <td>{process.Name}</td>
                                <td>{process.CPUPercent}
                                    <br />
                                    <Sparklines data={cpuHistory[process.PID]} limit={5} width={100} height={20} margin={5}>
                                        <SparklinesLine color="blue" />
                                    </Sparklines>
                                </td>
                                <td>{process.MemoryUsage}
                                    <br />
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
}

export default TaskManager;