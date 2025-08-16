import { articlePage } from './scenes/article_page/page'
import NewspaperLikeSite from './scenes/landing_page/landing_page'
import Header from './components/header/header'
import './App.css'
function App() {
  
  return (
    <>
      <Header/>
      <div className="m-20">
        {/* <NewspaperLikeSite/>       */}
        <articlePage title="article page" date="22/01/2004" description= link={} keywords={}/>
      </div>
    </>
  )
}

export default App
