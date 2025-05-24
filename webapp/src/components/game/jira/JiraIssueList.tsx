import React, { useEffect, useState } from 'react';
import { Box, Center, Text, Button } from "@chakra-ui/react";

interface Issue {
    key: string;
    summary: string;
}

const mockIssues: Issue[] = [
    { key: 'JRA-101', summary: 'Implementar autenticação de usuário' },
    { key: 'JRA-102', summary: 'Criar tela de login' },
    { key: 'JRA-103', summary: 'Desenvolver dashboard principal' },
];

const JiraIssueList: React.FC = () => {
    const [issues, setIssues] = useState<Issue[]>([]);

    useEffect(() => {
        // Simula uma chamada à API do Jira
        setTimeout(() => {
            setIssues(mockIssues);
        }, 1000);
    }, []);

    return (
        <div style={{ position: 'absolute', marginRight: 10}}>
            <h3>Histórias do Jira</h3>
            <ul>
                {issues.map((issue) => (

                    <Box
                        bg="blue.light"
                        w={{ base: "50vw", md: "30vw", lg: "20vw" }}
                        h={{ base: "14vh", md: "17vh", lg: "20vh" }}
                        transition="0.2s ease"
                        p={4}
                        borderRadius="3xl"
                        alignItems="center"
                        key={issue.key}
                    >
                        <Text fontSize="xl" textAlign="center" color="gray.600">
                            {issue.key}: {issue.summary}
                        </Text>
                    </Box>

                ))}
            </ul>
        </div>
    );
};

export default JiraIssueList;
