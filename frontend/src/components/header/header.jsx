import { useState, useEffect } from "react";

import { Popover} from "@/components/retroui/Popover";
import { Label} from "@/components/retroui/Label";
import { Button} from "@/components/retroui/Button";
import { Input } from "@/components/retroui/Input";

function PopoverStyleDefault() {
    let [url, SetUrl] = useState("");
    let [fetchedData, setFetchedData] = useState(null);
    console.log("URL", url );


    return (
        <Popover>
            <Popover.Trigger asChild>
                <Button>Scrape</Button>
            </Popover.Trigger>
            <Popover.Content className="w-80 font-sans">
                <Label className="mb-10">Enter URL</Label>


                <Input className="mb-5" placeholder ="Whatever you want" onChange={(e)=>{SetUrl(e.target.value)}}>
                   
                </Input>
                <Button>Scrape</Button>

            </Popover.Content>
        </Popover>
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
                    <><PopoverStyleDefault /></>
                    <div className="hover:underline decoration-primary mt-2 "><a href="#pricing">Pricing</a></div>
                    <div className="hover:underline decoration-primary mt-2 "> <a href="#contact">Contact</a> </div>
                </nav>
            </div>
        </header>
    );
}


