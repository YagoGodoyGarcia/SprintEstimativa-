import { Link as ChakraLink, Text, Button } from "@chakra-ui/react";
import Link from "next/link";
import { Hero } from "../components/general/Hero";
import { Container } from "../components/general/Container";
import { DarkModeSwitch } from "../components/general/DarkModeSwitch";
import { Footer } from "../components/general/Footer";

const Index = () => {
  return (
    <Container height="100vh" justifyContent="center">
      <DarkModeSwitch />
      <Hero>
        <Link href="/newgame">
          <Button variant="solid" colorScheme="red" size="lg">
            Iniciar novo jogo
          </Button>
        </Link>
      </Hero>
      <Footer>

      </Footer>
    </Container>
  );
};

export default Index;
