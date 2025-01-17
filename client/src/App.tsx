import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { DashboardClient } from './protos/dashboard.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Greeting } from './protos/dashboard'

function App() {
  const [count, setCount] = useState(0)
  const [greeting, setGreeting] = useState<Greeting | undefined>(undefined);

  const transport = new GrpcWebFetchTransport({
    baseUrl: "http://localhost:8070",
    format: 'binary'
  });
  const client = new DashboardClient(transport);

  useEffect(() => {
    client.getGreeting({}).then(({ response }) => {
      setGreeting(response.greeting);
    }).catch((err) => {
      console.log(err);
    })
  }, [])


  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      {greeting && (
        <div>
          <pre>{greeting.id}</pre>
          <p>{greeting.message}</p>
        </div>
      )}
    </>
  )
}

export default App
