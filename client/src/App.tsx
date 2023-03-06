import './App.css'
import React, {FC, FormEventHandler, useCallback, useState} from "react";
import useShortUrl from "./hooks/useShortUrl";
import CopiableLink from "./components/CopiableLink";

const App: FC = () => {
  const [url, setUrl] = useState('')

  const {shortUrl, fetchShortUrl} = useShortUrl()

  const onSubmit = useCallback<FormEventHandler<HTMLFormElement>>((e) => {
    e.preventDefault()

    setUrl('')
    fetchShortUrl(url)
  }, [url])

  return (
    <div className="App">
      <h1>Input your URL</h1>
      <form onSubmit={onSubmit}>
        <input type="text" value={url} onChange={e => setUrl(e.target.value)}/>
        <button type='submit'>Submit</button>
      </form>

      {shortUrl && <div>
          <h2>Your Short URL:</h2>
          <CopiableLink url={shortUrl}/>
      </div>}
    </div>
  )
}

export default App
