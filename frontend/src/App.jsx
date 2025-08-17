import { NewsPage } from './scenes/article_page/page'
import NewspaperLikeSite from './scenes/landing_page/landing_page'
import Header from './components/header/header'
import './App.css'
import { createBrowserRouter, Outlet, RouterProvider } from 'react-router-dom'
import { ArticleLoader } from './loader/article_loader'



function Layout(){
  return (
    <div className="flex flex-col space-y-10">
      <div>
        <Header />
      </div>
      <div className='max-w-7xl mx-auto'>
        <Outlet />
      </div>
    </div>
  );



}


const router = createBrowserRouter([
  {

    path:"/",
    element:<Layout/>,
    children:[
      {
        path:"",
        element:<NewspaperLikeSite/>
      },
      {
        path: "articles/:url",
        element: <NewsPage />,
        loader: ArticleLoader
      }
      
    ]
    
  },
  ])







function App() {
  
  return (
    <>
      
      
        {/* <NewspaperLikeSite/>       */}
        <RouterProvider router={router}/>
        {/* <NewsPage 
        title="article page" date="22/01/2004" description={str} link1={"link"} keywords={["keywords"]}/>
         */}

    </>
  )
}

export default App
