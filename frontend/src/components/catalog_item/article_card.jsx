// eslint-disable-next-line no-unused-vars
import { motion } from "framer-motion";
import {Card}  from "@/components/retroui/Card";


export function ArticleCard({ title ="This is Card Title", text}) {
    return (
        <div>

        <motion.div whileHover={{scale:1.1}}>
                <Card>
                    <Card.Header>
                        <Card.Title>{title}</Card.Title>
                        <Card.Description>
                            {text}
                        </Card.Description>
                    </Card.Header>
                </Card>
        </motion.div>
        
        </div>
    );
}
