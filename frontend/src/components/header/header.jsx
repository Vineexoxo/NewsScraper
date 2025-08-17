import { useState, useEffect } from "react";

import { Dialog} from "@/components/retroui/Dialog";
import { Label} from "@/components/retroui/Label";
import { Button} from "@/components/retroui/Button";
import { Input } from "@/components/retroui/Input";
import ProgressBar from "../progress_bar/progress_bar";
import { useFetch } from "@/hooks/useFetch";
import { Text } from "@/components/retroui/Text";
function DialogStyleDefault() {
    let [url, SetUrl] = useState("");
    let [clicked, setClicked] = useState(false);
    const { data, loading, error, fetchData } = useFetch(
        '/api/v1/article_by_url/',
        {
            "url": url,
        }
    );
    console.log("URL", url );


    return (
        <Dialog>
            <Dialog.Trigger asChild>
                <Button>Scrape</Button>
            </Dialog.Trigger>
            <Dialog.Content size={"md"}>
                    <Dialog.Header position={"fixed"}>
                        <Text as="h5">Scrape the Article</Text>
                    </Dialog.Header>
                {error!=null?
                    (<form className="flex flex-col gap-4">
                        <p className="text-red-500">The Url you have Entered is probably not a valid url</p>:
                    </form>):
                clicked ? (
                    <form className="flex flex-col gap-4">
                        <ProgressBar loading={loading} data={data}/>
                    </form>
                 ) :(


                        <form className="flex flex-col gap-2">
                                <div className="flex flex-col p-4 gap-2">
                                <Label className="mb-10">Enter URL</Label>
                                <Input className="mb-5" placeholder="Whatever you want"
                                    onChange={(e) => { SetUrl(e.target.value) }}>

                                </Input>
                                </div>
                                <Dialog.Footer>
                                    
                                        <Button onClick={() => {
                                            fetchData();
                                            setClicked(true);
                                        }}>
                                            Scrape
                                        </Button>
                                
                                </Dialog.Footer>
                        </form>      
                )}
            
            </Dialog.Content>
            
        </Dialog>
    );
}


export default function Header() {
    const [scrolled, setScrolled] = useState(false);
    
    useEffect(() => {
        const onScroll = () => setScrolled(window.scrollY > 50);
        window.addEventListener("scroll", onScroll);
        return () => window.removeEventListener("scroll", onScroll);
    }, []);

    return (
        <header
            className={`fixed top-0 left-0 w-full z-50 transition-all duration-300 ${scrolled
                    ? "bg-white/80 backdrop-blur-md shadow-md"
                    : "bg-transparent"
                }`}
        >
            <div className="max-w-7xl mx-auto px-4 py-2 flex items-center justify-between">
                <h1 className="text-black font-bold text-lg">NewsScraper</h1>
                <nav className="flex space-x-6 text-black">
                    <><DialogStyleDefault/></>
                    <div className="hover:underline decoration-primary mt-2 "><a href="#pricing">Pricing</a></div>
                    <div className="hover:underline decoration-primary mt-2 "> <a href="#contact">Contact</a> </div>
                </nav>
            </div>
        </header>
    );
}


