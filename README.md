Learning how to write parsers by implementing an expression evaluator.
Uses recursive-descent parsing technique.
Supports variables.

Grammar:
```
Statement :=
    | ID "=" Expression
    | Expression
    ;

Expression :=
    | Term ( ( "+" | "-" ) Term )*
    ;

Term := 
    | Factor ( ( "*" | "/" ) Factor )*;
    ;

Factor :=
    | NUMBER
    | ID
    | "-" Factor
    | "(" Expression ")"
    ;
```

TODO:
- Add more operators ( ^ % |abs| ! )
- Add global constants
